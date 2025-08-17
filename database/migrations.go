package database

import (
    students "student-portal/models/students"
)

func Migrate() error {
    return DB.AutoMigrate(
        &students.Student{},
        &students.Result{},
    )
}
