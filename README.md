# myerrors

```
type target struct {
    Msg string
    ...
    ...
    ...
}

func (t *target) Error() string {
    return t.Msg
}

err, ok := myerror.As[*target](err)


```