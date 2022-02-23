package structuredlog

import (
	"context"
	"fmt"
)

type OptionBuilder struct {
	options []Option
}

type Option func(e *Event)

func (ob *OptionBuilder) With(o Option)  *OptionBuilder {
	ob.options = append(ob.options, o)
	return ob
}

func (ob *OptionBuilder) WithVisibility(vis bool) *OptionBuilder {
	ob.options = append(ob.options, func(e *Event) {
		if !vis {
			var visInt = 0
			e.Visibility = &visInt
		}
	})
	return ob
}

func (ob *OptionBuilder) WithName(name string) *OptionBuilder {
	ob.options = append(ob.options, func(e *Event) {
		e.Name = name
	})
	return ob
}

func (ob *OptionBuilder) WithHttp(http Http) *OptionBuilder {
	ob.options = append(ob.options, func(e *Event) {
		if e.Attributes == nil {
			e.Attributes = &Attributes{}
		}
		e.Attributes.Http = &http
	})
	return ob
}

func (ob *OptionBuilder) WithDB(db DB) *OptionBuilder {
	ob.options = append(ob.options, func(e *Event) {
		if e.Attributes == nil {
			e.Attributes = &Attributes{}
		}
		e.Attributes.DB = &db
	})
	return ob
}

func (ob *OptionBuilder) WithException(err Exception)  *OptionBuilder {
	ob.options = append(ob.options, func(e *Event) {
		if e.Attributes == nil {
			e.Attributes = &Attributes{}
		}
		e.Attributes.Exception = &err
	})
	return ob
}

func With(o Option) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.With(o)
	return ob
}

func WithVisibility(vis bool) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.WithVisibility(vis)
	return ob
}

func WithName(name string) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.WithName(name)
	return ob
}

func WithHttp(http Http) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.WithHttp(http)
	return ob
}

func WithDB(db DB) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.WithDB(db)
	return ob
}

func WithException(err Exception) *OptionBuilder {
	ob := &OptionBuilder{}
	ob.WithException(err)
	return ob
}

func (ob *OptionBuilder) Debug(ctx context.Context, v ...interface{}) {
	std.output(ctx, SeverityDebug, fmt.Sprint(v...), ob.options)
}

func (ob *OptionBuilder) Info(ctx context.Context, v ...interface{}) {
	std.output(ctx, SeverityInfo, fmt.Sprint(v...), ob.options)
}

func (ob *OptionBuilder) Error(ctx context.Context, v ...interface{}) {
	std.output(ctx, SeverityError, fmt.Sprint(v...), ob.options)
}

func (ob *OptionBuilder) Debugf(ctx context.Context, format string, v ...interface{}) {
	std.output(ctx, SeverityDebug, fmt.Sprintf(format, v...), ob.options)
}

func (ob *OptionBuilder) Infof(ctx context.Context, format string, v ...interface{}) {
	std.output(ctx, SeverityInfo, fmt.Sprintf(format, v...), ob.options)
}

func (ob *OptionBuilder) Errorf(ctx context.Context, format string, v ...interface{}) {
	std.output(ctx, SeverityError, fmt.Sprintf(format, v...), ob.options)
}
