package singlylinkedlists

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSinglylinkedList_GetFirstValue(t *testing.T) {

	list := New()

	_, err := list.GetFirstValue()
	assert.NotNil(t, err)

	list.Add("abc")

	value, err := list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.Add("def")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.InsertAt(0, "xyz")

	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")
}

func TestSinglylinkedList_GetLastValue(t *testing.T) {
	list := New()

	_, err := list.GetLastValue()
	assert.NotNil(t, err)

	list.Add("abc")

	value, err := list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")

	list.Add("def")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "def")

	list.InsertAt(2, "xyz")

	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")
}

func TestSinglylinkedList_GetValues(t *testing.T) {
	list := New()

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))

	list.Add("a", "b", "c")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "b", "c"}))

	list.InsertAt(1, "d")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{"a", "d", "b", "c"}))

}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	list := New()
	assert.Equal(t, list.IsEmpty(), true)

	list.Add("normal_http_call")
	assert.Equal(t, list.IsEmpty(), false)

	list.RemoveAt(0)
	assert.Equal(t, list.IsEmpty(), true)
}

func TestSinglylinkedList_Size(t *testing.T) {
	list := New()
	assert.Equal(t, list.Size(), 0)

	list.Add("normal_http_call")
	assert.Equal(t, list.Size(), 1)

	list.Add("abc")
	assert.Equal(t, list.Size(), 2)

	list.RemoveAt(0)
	assert.Equal(t, list.Size(), 1)

	list.RemoveAt(0)
	assert.Equal(t, list.Size(), 0)
}

func TestSinglylinkedList_GetIndexOf(t *testing.T) {
	list := New()

	assert.Equal(t, list.GetIndexOf("normal_http_call"), -1)

	list.Add("normal_http_call")
	assert.Equal(t, list.GetIndexOf("normal_http_call"), 0)

	list.Add("abc")
	assert.Equal(t, list.GetIndexOf("abc"), 1)

	list.Add("abc")
	assert.Equal(t, list.GetIndexOf("abc"), 1)
}

func TestSinglylinkedList_InsertAt(t *testing.T) {
	list := New()

	err := list.InsertAt(3, "normal_http_call")
	assert.NotNil(t, err)

	err = list.InsertAt(0, "normal_http_call")
	assert.Nil(t, err)
}

func TestSinglylinkedList_RemoveAt(t *testing.T) {
	list := New()

	_, err := list.RemoveAt(0)
	assert.NotNil(t, err)

	list.Add("normal_http_call")
	value, err := list.RemoveAt(0)
	assert.Equal(t, value, "normal_http_call")
	assert.Nil(t, err)
}

func TestSinglylinkedList_Clear(t *testing.T) {

	list := New()

	list.Clear()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)

	list.Add("normal_http_call")

	list.Clear()

	assert.True(t, reflect.DeepEqual(list.GetValues(), []interface{}{}))
	assert.Equal(t, list.IsEmpty(), true)
	assert.Equal(t, list.Size(), 0)
}
