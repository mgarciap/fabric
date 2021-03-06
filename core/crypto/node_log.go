/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crypto

func (node *nodeImpl) info(format string, args ...interface{}) {
	log.Infof(node.conf.logPrefix+format, args...)
}

func (node *nodeImpl) debug(format string, args ...interface{}) {
	log.Debugf(node.conf.logPrefix+format, args...)
}

func (node *nodeImpl) error(format string, args ...interface{}) {
	log.Errorf(node.conf.logPrefix+format, args...)
}

func (node *nodeImpl) warning(format string, args ...interface{}) {
	log.Warningf(node.conf.logPrefix+format, args...)
}
