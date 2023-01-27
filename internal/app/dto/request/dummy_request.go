package request

type DummyRequest struct {
	Data interface{} `json:"data" validate:"required"`
}
