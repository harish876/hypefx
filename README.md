# Usage
1. Install the hypefx release build using `go install github.com/harish876/hypefx@latest`.
2. Install the hypefx test build using `go install --tags=test github.com/harish876/hypefx@latest`
3. Run `hypefx generate [project_name]` to scaffold a project. This initializes a go module with the project name and sets default configs at "hypefxconfig.json".
4. Create a routes folder in your root directory and Use `hypefx build` to build the routes from HTTP handlers defined in the app directory. This is your automagic file based routing.
5. Add the `RegisterRoutes` method to your "cmd/main.go".
6. Use `npm run dev` to start the development server.

# Future Plans
1. I have a lot of cool stuff planned for this, I want to remove many small bugs in the code and improve code quality.
2. Make nightly builds and automate a few of these manual processes listed under the "Usage" section.
3. Improving file-based routing by baking it into the dev server, removing multiple dependency commands, and integrating them into one.
4. Add a ton of batteries for common utilities found in web apps and things I use in my day-to-day life.

# Issue or Errors.
1. If you encounter an error, please file an issue here at github.com/harish876/hypefx or even better a pull request if you can diagnose and fix the error. Check out the test build for debug logs and off course the source code here.
