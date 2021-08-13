/*
 *
 * Copyright 2021 waterdrop authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCbc(t *testing.T) {
	content := []byte("Hello Waterdrop!")
	cbc, err := NewAesCbcCrypto("abcdefghijklmnop")
	assert.Equal(t, nil, err)
	encypt, err := cbc.Encrypt(content)
	assert.Nil(t, err)
	decrypt, err := cbc.Decrypt(encypt)
	assert.Nil(t, err)
	assert.Equal(t, content, decrypt)
	baseEncrypt, err := cbc.EncryptToString(content, Base64)
	assert.Nil(t, err)
	dst, err := cbc.DecryptFromString(baseEncrypt, Base64)
	assert.Equal(t, nil, err)
	assert.Equal(t, content, dst)
}
