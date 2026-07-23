package dblib:
import (
	 "database/sql"
	 _ "github.com/lib/pq" 
)
type EmpDAO struct {
}
var dsn = "postgres://postgres:MyPassword!23@35.226.182.49/postgres?sslmode=disable"
   
func (e *EmpDAO) List() ([]Emp, error) {
	 db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("select * from emp")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var emps []Emp

    for rows.Next() {
        emp := Emp{}
        err := rows.Scan(&emp.Empno, &emp.Ename, &emp.Salary)
        if err != nil {
            return nil, err
        }
        emps = append(emps, emp)
    }
    return emps, nil

}


func (e *EmpDAO) Insert(emp *Emp) error {
    //open connection
   
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    //insert record
    row, err := db.Exec("insert into emp(empno, ename, salary) values($1, $2, $3)", emp.Empno, emp.Ename, emp.Salary)
    if err != nil {
        return err
    }
    _, err = row.RowsAffected(); 
    return  err	
}
func (e *EmpDAO) Update(emp *Emp) error {
    //open connection
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    //update record
    row, err := db.Exec("update emp set ename=$1, salary=$2 where empno=$3", emp.Ename, emp.Salary, emp.Empno)
    if err != nil {
        return err
    }
    _, err = row.RowsAffected()
    return err
}
func (e *EmpDAO) Delete(empno int32) error {
    //open connection
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    //update record
    row, err := db.Exec("delete from emp where empno=$1", empno)
    if err != nil {
        return err
    }
    _, err = row.RowsAffected()
    return err
}



