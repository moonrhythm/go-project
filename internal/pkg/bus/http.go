package bus

import (
	"encoding/json"
	"errors"
	"mime"
	"net/http"

	"github.com/moonrhythm/dispatcher"

	"go-project/internal/pkg/httperror"
)

var wrapper = dispatcher.HTTPHandlerWrapper{
	Result:     "Result",
	Dispatcher: mux,
	Encoder:    httpEncoder,
	Decoder:    httpDecoder,
}

// HTTPHandler wraps message into http handler
func HTTPHandler(msg Message) http.Handler {
	return wrapper.Handler(msg)
}

func httpEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	statusCode := http.StatusOK

	if err, ok := v.(error); ok {
		httpErr := httperror.From(err)
		statusCode = httpErr.StatusCode()
		v = httpErr
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

type validator interface {
	Valid() error
}

var (
	errMethodNotAllowed     = errors.New("bus: method not allowed")
	errUnsupportedMediaType = errors.New("bus: unsupported media type")
)

func init() {
	httperror.Register(errMethodNotAllowed, httperror.New(http.StatusMethodNotAllowed, "Method Not Allowed"))
	httperror.Register(errUnsupportedMediaType, httperror.New(http.StatusUnsupportedMediaType, "Unsupported Media Type"))
}

func httpDecoder(r *http.Request, v interface{}) error {
	if r.Method != http.MethodPost {
		return errMethodNotAllowed
	}

	mt, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return errUnsupportedMediaType
	}
	switch mt {
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			return err
		}
	default:
		return errUnsupportedMediaType
	}

	if v, ok := v.(validator); ok {
		err = v.Valid()
		if err != nil {
			return err
		}
	}

	return nil
}
