package database

import (
    students "student-portal/models/students" // ← give alias
)

func Migrate() {
    err := DB.AutoMigrate(
        &students.Student{},
        &students.Result{},
    )

    if err != nil {
        panic("Failed to run migrations: " + err.Error())
    }
}
