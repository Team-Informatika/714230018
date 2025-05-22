package main

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type Staff struct {
	Name     string
	Position string
	Email    string
	Phone    string
}

func StaffForm(c *fiber.Ctx) error {
	formHTML := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Staff Data Form</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				form { max-width: 400px; margin: auto; }
				label { display: block; margin-top: 10px; }
				input[type="text"], input[type="email"], input[type="tel"] {
					width: 100%; padding: 8px; box-sizing: border-box;
				}
				button {
					margin-top: 15px; padding: 10px 20px; background-color: #4CAF50; color: white; border: none; cursor: pointer;
				}
				button:hover {
					background-color: #45a049;
				}
			</style>
		</head>
		<body>
			<h2>Staff Data Form</h2>
			<form action="/staffform" method="POST">
				<label for="name">Name:</label>
				<input type="text" id="name" name="name" required>

				<label for="position">Position:</label>
				<input type="text" id="position" name="position" required>

				<label for="email">Email:</label>
				<input type="email" id="email" name="email" required>

				<label for="phone">Phone:</label>
				<input type="tel" id="phone" name="phone" required>

				<button type="submit">Submit</button>
			</form>
		</body>
		</html>
	`
	return c.Type("html").SendString(formHTML)
}

func StaffFormSubmit(c *fiber.Ctx) error {
	staff := Staff{
		Name:     c.FormValue("name"),
		Position: c.FormValue("position"),
		Email:    c.FormValue("email"),
		Phone:    c.FormValue("phone"),
	}

	const resultTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Staff Data Submitted</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			.container { max-width: 500px; margin: auto; padding: 20px; border: 1px solid #ddd; border-radius: 5px; background-color: #f9f9f9; }
			h2 { color: #4CAF50; }
			p { font-size: 18px; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Staff Data Submitted</h2>
			<p><strong>Name:</strong> {{.Name}}</p>
			<p><strong>Position:</strong> {{.Position}}</p>
			<p><strong>Email:</strong> {{.Email}}</p>
			<p><strong>Phone:</strong> {{.Phone}}</p>
			<a href="/staffform">Back to form</a>
		</div>
	</body>
	</html>
	`

	tmpl, err := template.New("result").Parse(resultTemplate)
	if err != nil {
		return c.Status(500).SendString("Template parsing error")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, staff); err != nil {
		return c.Status(500).SendString("Template execution error")
	}

	return c.Type("html").SendString(buf.String())
}
