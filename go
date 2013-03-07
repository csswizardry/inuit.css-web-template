#!/bin/sh

echo "Installing clean project"

mv .git/modules .       && \
rm -rf .git/            && \
git init                && \
mv modules .git/        && \
git add css/inuit.css   && \
git add .gitmodules     && \
rm $0

echo "inuit.css project successfully installed"

git status
