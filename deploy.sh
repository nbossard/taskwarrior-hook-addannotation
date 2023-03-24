#! /bin/sh
#
# deploy.sh
# Copyright (C) 2023 nbossard <nbossard@MC-C02Y91ZDJGH5>
#
# Distributed under terms of the MIT license.
#

go build addissueannotation
cp  addissueannotation ~/.task/hooks/on-add-autoaddannotations

