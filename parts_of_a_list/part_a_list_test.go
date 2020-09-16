package parts_of_a_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartList(t *testing.T) {

	assert.Equal(t, PartList([]string{"I", "wish", "I", "hadn't", "come"}), "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")

	assert.Equal(t, PartList([]string{"cdIw", "tzIy", "xDu", "rThG"}), "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)")

}
