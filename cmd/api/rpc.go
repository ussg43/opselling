package main

import "errors"

type Args struct {
	sheetID string
}

func (a *Application) UpdateSheet(args *Args, reply *bool) error {
	if args == nil {
		return errors.New("arguments cannot be empty")
	}

	

	*reply = true
	return nil
}