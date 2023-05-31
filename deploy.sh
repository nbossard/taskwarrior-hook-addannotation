#! /bin/sh
#
# deploy.sh
# Copyright (C) 2023 nbossard <nbossard@MC-C02Y91ZDJGH5>
#
# Distributed under terms of the MIT license.
#

echo "building hookaddannotation..."
go build hookaddannotation
echo "copying hookaddannotation to ~/.task/hooks"
cp  hookaddannotation ~/.task/hooks/on-add-autoaddannotations
cp  hookaddannotation ~/.task/hooks/on-modify-autoaddannotations
echo "done"
