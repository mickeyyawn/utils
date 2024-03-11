

	err := jsonparser.ObjectEach(*e.buf, func(key, value []byte, vt jsonparser.ValueType, offset int) error {

		//var temp

		switch string(key) {
		case zerolog.MessageFieldName:
			temp := bytesToStrUnsafe(value)
		case zerolog.ErrorFieldName:
			/*
			event.Exception = append(event.Exception, sentry.Exception{
				Value:      bytesToStrUnsafe(value),
				Stacktrace: newStacktrace(),
			})*/
		case zerolog.LevelFieldName, zerolog.TimestampFieldName:
		default:
			//event.Extra[string(key)] = bytesToStrUnsafe(value)
		}

		//return nil
	})


	if err != nil {
		return
	}
