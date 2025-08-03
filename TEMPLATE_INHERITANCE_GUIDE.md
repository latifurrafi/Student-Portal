# Go Template Inheritance Guide

## ✅ **Yes, Template Inheritance is Possible in Go!**

Go templates support inheritance through the `{{template}}` and `{{define}}` directives. Here's how it works:

## 🏗️ **How Go Template Inheritance Works**

### 1. **Base Layout Template** (`views/layout.html`)
```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    {{template "head" .}}
</head>
<body>
    <header>{{template "header" .}}</header>
    <main>{{template "content" .}}</main>
    <footer>{{template "footer" .}}</footer>
    {{template "scripts" .}}
</body>
</html>
```

### 2. **Child Template** (`views/dashboard_inherited.html`)
```html
{{template "layout" .}}

{{define "head"}}
<!-- Page-specific head content -->
<style>.dashboard-specific { color: blue; }</style>
{{end}}

{{define "content"}}
<!-- Page-specific content -->
<h1>Dashboard Content</h1>
{{end}}

{{define "scripts"}}
<!-- Page-specific scripts -->
<script>console.log('Dashboard loaded');</script>
{{end}}
```

## 🔧 **Key Concepts**

### **Template Blocks**
- `{{define "blockName"}}` - Defines a template block
- `{{template "blockName" .}}` - Includes a template block
- `{{template "layout" .}}` - Includes the base layout

### **Data Passing**
- The `.` (dot) passes the current data context
- All blocks receive the same data structure

## 📁 **File Structure**
```
views/
├── layout.html              # Base layout template
├── dashboard_inherited.html # Inherits from layout
├── profile_inherited.html   # Inherits from layout
├── dashboard_standalone.html # Self-contained (no inheritance)
└── profile_standalone.html  # Self-contained (no inheritance)
```

## 🚀 **Testing Both Approaches**

### **Standalone Templates (No Inheritance)**
- `/dashboard` - Uses `dashboard_standalone.html`
- `/profile` - Uses `profile_standalone.html`

### **Inherited Templates (With Inheritance)**
- `/dashboard-inherited` - Uses `dashboard_inherited.html`
- `/profile-inherited` - Uses `profile_inherited.html`

## 💡 **Advantages of Template Inheritance**

### ✅ **Pros:**
- **DRY Principle** - Don't repeat common layout code
- **Consistency** - All pages share the same base structure
- **Maintainability** - Change layout once, affects all pages
- **Modularity** - Each page only defines its unique content

### ⚠️ **Cons:**
- **Complexity** - Can be harder to debug
- **File Dependencies** - Templates depend on each other
- **Caching Issues** - Template engine might cache incorrectly

## 🔍 **Common Issues & Solutions**

### **Issue 1: Template Not Found**
```go
// ❌ Wrong - template name mismatch
return c.Render("dashboard", data)

// ✅ Correct - ensure template exists
return c.Render("dashboard_inherited", data)
```

### **Issue 2: Missing Block Definition**
```html
<!-- ❌ Missing block definition -->
{{template "content" .}}

<!-- ✅ Define the block -->
{{define "content"}}
<h1>Content here</h1>
{{end}}
```

### **Issue 3: Data Context Issues**
```go
// ❌ Wrong - missing data
return c.Render("template", nil)

// ✅ Correct - pass required data
return c.Render("template", fiber.Map{
    "Title": "Page Title",
    "Data": someData,
})
```

## 🎯 **Best Practices**

### **1. Use Clear Naming**
```html
<!-- Good -->
{{define "page_head"}}
{{define "page_content"}}
{{define "page_scripts"}}

<!-- Avoid -->
{{define "h"}}
{{define "c"}}
{{define "s"}}
```

### **2. Provide Default Blocks**
```html
<!-- In layout.html -->
{{define "head"}}{{end}}
{{define "scripts"}}{{end}}
{{define "content"}}<p>Default content</p>{{end}}
```

### **3. Use Consistent Data Structure**
```go
type PageData struct {
    Title      string
    ActivePage string
    User       User
    // ... other common fields
}
```

### **4. Handle Missing Data Gracefully**
```html
{{if .User}}
    <p>Welcome, {{.User.Name}}</p>
{{else}}
    <p>Welcome, Guest</p>
{{end}}
```

## 🔄 **Template Engine Configuration**

```go
engine := html.New("./views", ".html")

// Enable template reloading in development
if os.Getenv("ENV") == "development" {
    engine.Reload(true)
}

app := fiber.New(fiber.Config{
    Views: engine,
})
```

## 📊 **Comparison: Inheritance vs Standalone**

| Aspect | Inheritance | Standalone |
|--------|-------------|------------|
| **Code Reuse** | ✅ High | ❌ Low |
| **Maintenance** | ✅ Easy | ⚠️ Harder |
| **Debugging** | ⚠️ Complex | ✅ Simple |
| **Performance** | ✅ Good | ✅ Good |
| **Flexibility** | ⚠️ Limited | ✅ High |

## 🎉 **Conclusion**

**Template inheritance in Go is not only possible but also powerful!** 

- **Use inheritance** when you have consistent layouts and want to maintain DRY principles
- **Use standalone templates** when you need maximum flexibility or are debugging template issues
- **Both approaches work well** - choose based on your project needs

The key is understanding how `{{define}}` and `{{template}}` work together to create reusable, maintainable templates. 