#!/bin/sh

echo "Installing clean project"

mv .git/modules .       && \
rm -rf .git/            && \
git init                && \
mv modules .git/        && \
git add css/inuit.css   && \
git add .gitmodules     && \
rm go

echo "inuit.css project successfully installed"

git status
