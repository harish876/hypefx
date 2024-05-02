# Usage - (**Incomplete*)

1. This is a common Grid component which can be reused across different grid/table views.
2. This can be customized and made your own by adding different styling or a different UI flavour.
3. Things to be added to start developing a grid endpoint.

    -  Implement your backend API routes like this. These endpoints need to be developed using the 
       default implementation of the grid.

       ```go
        func RegisterGridRoutes(e *echo.Echo) {
	        apiGroup := e.Group("/grid")

	        apiGroup.GET("", handlers.GridHandler)
	        apiGroup.POST("", handlers.GridFilterHandler)

	        apiGroup.GET("/:id", handlers.GridRowHandler)
	        apiGroup.PUT("/:id", handlers.UpdateGridRowHandler)
	        apiGroup.DELETE("/:id", handlers.DeleteGridRowHandler)

	        apiGroup.GET("/edit/:id", handlers.RenderEditGridHandler)
        }
       ```


    - Implement your handler functions corresponding to these routes. Here is what a grid config and 
    a grid model looks like. For more info, checkout the code snippets in the examples folder.

    ```go
        var columns = []grid.GridColumn{
        {
            Typ:      grid.String,
            Label:    "Name",
            Key:      "Name",
            Renderer: "name",
        },
        {
            Typ:      grid.String,
            Label:    "Status",
            Key:      "Status",
            Renderer: "status",
            Editable: true,
            EditOptions: grid.GridEditOptions{
                EditType: grid.EditSelect,
                EditProps: grid.GridSelectEditProps{
                    Id:    "Id",
                    Name:  "status",
                    Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
                    Options: []dropdown.SelectOption{
                        {Label: "Active", Value: "active"},
                        {Label: "Inactive", Value: "inactive"},
                    },
                },
            },
        },
        {
            Typ:      grid.String,
            Label:    "Role",
            Key:      "Position",
            Editable: true,
            EditOptions: grid.GridEditOptions{
                EditType: grid.EditInput,
                EditProps: grid.GridInputEditProps{
                    Id:    "Id",
                    Typ:   input.InputTypeText,
                    Name:  "position",
                    Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
                },
            },
        },
        {
            Typ:      grid.Array,
            Label:    "Badges",
            Key:      "Badges",
            Renderer: "badges",
            Editable: true,
            EditOptions: grid.GridEditOptions{
                EditType: grid.EditMultiSelect,
                EditProps: grid.GridMultiSelectEditProps{
                    Id:   "Id",
                    Name: "badges",
                    Options: []dropdown.SelectOption{
                        {Label: "Design", Value: "Design"},
                        {Label: "Product", Value: "Product"},
                        {Label: "Marketing", Value: "Marketing"},
                        {Label: "Engineering", Value: "Engineering"},
                        {Label: "Analytics", Value: "Analytics"},
                        {Label: "UI Design", Value: "UI Design"},
                    },
                },
            },
        },
        {
            Typ:   grid.String,
            Label: "Email",
            Key:   "Email",
        },
        }

        var gridModel = grid.GridContext[grid.GridColumn]{
            Title:       "Customers",
            Subheading:  "",
            Description: "Example Gird with filters, pagination, export etc...",
            Columns:     columns,
            IdField:     "Id",
            Url:         "/grid",
        }

    ``` 