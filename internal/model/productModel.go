package model

import "errors"

type ProdutoDomain struct {
    ID          string
    UserID      string
    URL         string
    TargetPrice float64
    IsActive    bool
}

func (p *ProdutoDomain) Validate() error {
    if p.URL == "" {
        return errors.New("url is required")
    }
    if p.TargetPrice <= 0 {
        return errors.New("target price must be greater than zero")
    }
    return nil
}