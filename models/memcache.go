package models

import (
	"github.com/bradfitz/gomemcache/memcache"
	"reflect"
)

// cache Объект кэша memcached (http://localhost:11211)
var cache = memcache.New("localhost:11211")

// CacheData Функция кэширования данных с использованием memcached, где
// cacheKey - ключ кэша,
// ttl - время жизни записи в кэше,
// fn  - функция, получающая данные в случае отсутствия данных в кэше. Может быть какая угодно, но должна возвращать слайс []byte
func CacheData(cacheKey string, ttl int32, fn any) []byte {

	// возвращаемое значение
	var retValue []byte

	// Пытаемся получить данные из кэша по ключу
	item, err := cache.Get(cacheKey)

	if err != nil {
		// если возникла ошибка - значит данных в кэше нет, или возникли проблемы с подключением
		// в случае отсутствия данных в кэше, используем функцию fn для получения этих данных из БД

		// берем конкретную величину, возвращаемую функцией fn
		fnValue := reflect.ValueOf(fn)
		// создаем слайс аргументов для хранения параметров функции fn
		args := make([]reflect.Value, fnValue.Type().NumIn())

		// проверяем, что функция возвращает слайс байтов ([]byte)
		if fnValue.Type().NumOut() != 1 || fnValue.Type().Out(0).Kind() != reflect.Slice || fnValue.Type().Out(0).Elem().Kind() != reflect.Uint8 {
			panic("функция должна возвращать тип []byte")
		}

		// вызываем функцию (она возвращает слайс типа reflect)
		results := fnValue.Call(args)

		// первый и единственный элемент слайса результата содержит нужные нам данные
		// преобразуем нулевой элемент слайса результата в слайс байт
		retValue = results[0].Interface().([]byte)

		// помещаем найденное в кэш
		memcacheItem := memcache.Item{
			Key:        cacheKey,
			Expiration: ttl,
			Value:      retValue,
		}
		// помещаем в кэш
		err := cache.Set(&memcacheItem)
		if err != nil {
			panic(err)
		}
	} else {
		// если ошибки получения из кэша нет - берем его
		retValue = item.Value
	}

	return retValue
}
