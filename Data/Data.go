package Data

import (
	"errors"
	"sync"
)

var MirrorsList = []string{
	"http://ftp.am.debian.org/debian/",
	"http://ftp.au.debian.org/debian/",
	"http://ftp.at.debian.org/debian/",
	"http://ftp.by.debian.org/debian/",
	"http://ftp.be.debian.org/debian/",
	"http://ftp.br.debian.org/debian/",
	"http://ftp.bg.debian.org/debian/",
	"http://ftp.ca.debian.org/debian/",
	"http://ftp.cl.debian.org/debian/",
	"http://ftp2.cn.debian.org/debian/",
	"http://ftp.cn.debian.org/debian/",
	"http://ftp.hr.debian.org/debian/",
	"http://ftp.cz.debian.org/debian/",
	"http://ftp.dk.debian.org/debian/",
	"http://ftp.sv.debian.org/debian/",
	"http://ftp.ee.debian.org/debian/",
	"http://ftp.fr.debian.org/debian/",
	"http://ftp2.de.debian.org/debian/",
	"http://ftp.de.debian.org/debian/",
	"http://ftp.gr.debian.org/debian/",
	"http://ftp.hk.debian.org/debian/",
	"http://ftp.hu.debian.org/debian/",
	"http://ftp.is.debian.org/debian/",
	"http://ftp.it.debian.org/debian/",
	"http://ftp.jp.debian.org/debian/",
	"http://ftp.kr.debian.org/debian/",
	"http://ftp.lt.debian.org/debian/",
	"http://ftp.mx.debian.org/debian/",
	"http://ftp.md.debian.org/debian/",
	"http://ftp.nl.debian.org/debian/",
	"http://ftp.nc.debian.org/debian/",
	"http://ftp.nz.debian.org/debian/",
	"http://ftp.no.debian.org/debian/",
	"http://ftp.pl.debian.org/debian/",
	"http://ftp.pt.debian.org/debian/",
	"http://ftp.ro.debian.org/debian/",
	"http://ftp.ru.debian.org/debian/",
	"http://ftp.sg.debian.org/debian/",
	"http://ftp.sk.debian.org/debian/",
	"http://ftp.si.debian.org/debian/",
	"http://ftp.es.debian.org/debian/",
	"http://ftp.fi.debian.org/debian/",
	"http://ftp.se.debian.org/debian/",
	"http://ftp.ch.debian.org/debian/",
	"http://ftp.tw.debian.org/debian/",
	"http://ftp.tr.debian.org/debian/",
	"http://ftp.uk.debian.org/debian/",
	"http://ftp.us.debian.org/debian/",
}

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
	Salary int    `json:"salary"`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e Employee) Employee
	Delete(id int) error
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()

	e.Id = s.counter
	s.data[e.Id] = *e

	s.counter++
	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	if _, exist := s.data[id]; exist == false {
		errors.New("Employee not found")
	}
	return s.data[id], nil
}

func (s *MemoryStorage) Delete(id int) error {
	if _, exist := s.data[id]; exist == false {
		return errors.New("User doesn't delete")
	}
	return nil
}

func (s *MemoryStorage) Update(id int, e Employee) Employee {
	emp, exist := s.data[id]

	if exist == false {
		errors.New("Employee not found")
	}

	emp.Name = e.Name
	emp.Age = e.Age
	emp.Sex = e.Sex
	emp.Salary = e.Salary

	s.data[id] = emp
	return emp
}
