package template

type TemplateParams struct {
	BasePath            string //path to the directory from which the handlers will be intepreted. Default app
	BaseImportPath      string //github.com/harish876/hypefx/internal/generators/test/app - revisit
	TemplateName        string //routes
	DestinationDir      string //path to directory where all routes will be generated
	RouteDirPackageName string //revisit default:routes
}
