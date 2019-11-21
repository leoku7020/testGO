<html>
{{if .Data.Success}}
    <h1>Thanks for your message!</h1>
    <div>What you Type in is:</div>
    <div>Email:{{.Data.Email}}</div>
    <div>Subject:{{.Data.Subject}}</div>
    <div>Message:{{.Data.Message}}</div>
    <div>Data:{{.Data}}</div>
{{else}}
    <h1>Contact</h1>
    <form method="POST">
        <label>Email:</label><br />
        <input type="text" name="email"><br />
        <label>Subject:</label><br />
        <input type="text" name="subject"><br />
        <label>Message:</label><br />
        <textarea name="message"></textarea><br />
        <input type="submit">
    </form>
{{end}}
</html>