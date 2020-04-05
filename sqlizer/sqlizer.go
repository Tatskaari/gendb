package sqlizer

func combineArgs(args []interface{}, newArgs []interface{}) []interface{} {
	combinedArgs := args
	for _, arg := range newArgs {
		combinedArgs = append(combinedArgs, arg)
	}
	return combinedArgs
}

