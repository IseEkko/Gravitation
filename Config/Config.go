/***
creat file time:2023-4-26,
This page defines the file reading interface
*/
package Config

type Config interface {
	Load() error              //load file
	Scan(v interface{}) error //scan file
	GetValue(key string)      //get key value
}
