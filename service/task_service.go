package service

type TaskService interface {
	Add() error
	List() error
}
