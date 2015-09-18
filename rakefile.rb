task :run do
		sh  "go get"
		sh  "go build"
		sh  "./schnapi"
end
