package models

type Technology struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LogoUrl  string `json:"logoUrl"`
	Category string `json:"category"`
}

// Fake database
var Technologies = []Technology{
	{ID: "1", Name: "Angular", LogoUrl: "https://img.icons8.com/color/48/angularjs.png", Category: "Frontend"},
	{ID: "11", Name: "React", LogoUrl: "https://img.icons8.com/color/48/react-native.png", Category: "Frontend"},
	{ID: "2", Name: "Symfony", LogoUrl: "https://img.icons8.com/color/48/symfony.png", Category: "Backend"},
	{ID: "22", Name: "Node.js", LogoUrl: "https://img.icons8.com/fluency/48/node-js.png", Category: "Backend"},
	{ID: "21", Name: "APIPlatform", LogoUrl: "/assets/img/api-platform.svg", Category: "Backend"},
	{ID: "5", Name: "Docker", LogoUrl: "https://img.icons8.com/color/48/docker.png", Category: "Techno"},
	{ID: "6", Name: "Jenkins", LogoUrl: "https://img.icons8.com/color/48/jenkins.png", Category: "Techno"},
	{ID: "7", Name: "FrankenPHP", LogoUrl: "https://img.icons8.com/color/48/php.png", Category: "Techno"},
	{ID: "8", Name: "Kubernetes", LogoUrl: "https://img.icons8.com/color/48/kubernetes.png", Category: "Techno"},
	{ID: "9", Name: "GitHub Actions", LogoUrl: "https://img.icons8.com/color/48/github.png", Category: "Techno"},
}
