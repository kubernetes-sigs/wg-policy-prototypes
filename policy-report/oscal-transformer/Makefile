# -*- mode:makefile; coding:utf-8 -*-

# Copyright 2022 The CNCF Policy Working Group Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


TRESTLE_WORKSPACE = /tmp/trestle.workspace
TRESTLE_INSTALL = /tmp/venv.trestle
YAML_TO_OSCAL = k8s_to_oscal.py

.SILENT: info
info:
	echo "run: make oscal";

.SILENT: install
install:
	if [ ! -d venv.trestle ]; then \
		echo "=> create python virtual environment venv.trestle"; \
		python -m venv $(TRESTLE_INSTALL); \
		source $(TRESTLE_INSTALL)/bin/activate; \
		echo "=> install trestle with prereqs"; \
		python -m pip install -q --upgrade pip setuptools; \
		python -m pip install -q compliance-trestle==1.1.0; \
		echo "=> create and initialize trestle workspace"; \
		mkdir $(TRESTLE_WORKSPACE); \
		cd $(TRESTLE_WORKSPACE); \
		trestle init >> /dev/null; \
	else \
		echo "=> using python virtual environment $(TRESTLE_INSTALL)"; \
	fi

.SILENT: clean
clean:
	echo "=> remove python virtual environment";
	rm -fr $(TRESTLE_INSTALL)
	echo "=> remove trestle workspace";
	rm -fr $(TRESTLE_WORKSPACE)

.SILENT: oscal
oscal: install
	echo "=> transform samples from k8s yaml to OSCAL json"; \
	echo ">> source $(TRESTLE_INSTALL)/bin/activate"; \
	source $(TRESTLE_INSTALL)/bin/activate; \
	echo ">> python $(YAML_TO_OSCAL)-to-oscal.py --in ../samples --out $(TRESTLE_WORKSPACE)/assessment-results"; \
	python $(YAML_TO_OSCAL) --in ../samples --out $(TRESTLE_WORKSPACE)/assessment-results;

.SILENT: oscal.partial
oscal.partial: install
	echo "=> transform samples from k8s yaml to OSCAL json"; \
	echo ">> source $(TRESTLE_INSTALL)/bin/activate"; \
	source $(TRESTLE_INSTALL)/bin/activate; \
	echo ">> python $(YAML_TO_OSCAL) --in ../samples --out $(TRESTLE_WORKSPACE)/assessment-results --ar-type partial"; \
	python $(YAML_TO_OSCAL) --in ../samples --out $(TRESTLE_WORKSPACE)/assessment-results --ar-type partial;

.SILENT: help
help: install
	echo ">> python $(YAML_TO_OSCAL) -h"; \
	python $(YAML_TO_OSCAL) -h;

