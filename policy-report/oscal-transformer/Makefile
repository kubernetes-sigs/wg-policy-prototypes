# -*- mode:makefile; coding:utf-8 -*-

# Copyright 2022 Lou DeGenaro (degenaro@us.ibm.com).
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

.SILENT: info
info:
	echo "run: make oscal";

.SILENT: install
install:
	if [ ! -d venv.trestle ]; then \
		echo "=> create python virtual environment venv.trestle"; \
		python -m venv venv.trestle; \
		source venv.trestle/bin/activate; \
		echo "=> install trestle with prereqs"; \
		python -m pip install -q --upgrade pip setuptools; \
		python -m pip install -q compliance-trestle; \
		echo "=> create and initialize trestle workspace"; \
		mkdir trestle.workspace; \
		cd trestle.workspace; \
		trestle init >> /dev/null; \
	else \
		echo "=> using python virtual environment venv.trestle"; \
	fi

.SILENT: clean
clean:
	echo "=> remove python virtual environment";
	rm -fr venv.trestle
	echo "=> remove trestle workspace";
	rm -fr trestle.workspace

.SILENT: oscal
oscal: install
	echo "=> transform samples from k8s yaml to OSCAL json"; \
	echo ">> source venv.trestle/bin/activate"; \
	source venv.trestle/bin/activate; \
	echo ">> python k8s-to-oscal.py --in ../samples --out trestle.workspace/assessment-results"; \
	python k8s-to-oscal.py --in ../samples --out trestle.workspace/assessment-results;

.SILENT: oscal.partial
oscal.partial: install
	echo "=> transform samples from k8s yaml to OSCAL json"; \
	echo ">> source venv.trestle/bin/activate"; \
	source venv.trestle/bin/activate; \
	echo ">> python k8s-to-oscal.py --in ../samples --out trestle.workspace/assessment-results --ar-type partial"; \
	python k8s-to-oscal.py --in ../samples --out trestle.workspace/assessment-results --ar-type partial;

.SILENT: help
help: install
	echo ">> python k8s-to-oscal.py -h"; \
	python k8s-to-oscal.py -h;

