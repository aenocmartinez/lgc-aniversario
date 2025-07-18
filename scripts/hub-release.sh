#!/bin/bash
image_name="lgc-aniversario-container"
tag=$1
allowed_tags=("qa" "latest" "stable" "test")
if [[ -z "$tag" ]]; then
    echo "Please provide a tag, one of: ${allowed_tags[@]}"
    exit 1
fi
found=false
for allowed_tag in "${allowed_tags[@]}"; do
    if [[ "$tag" == "$allowed_tag" ]]; then
        found=true
        break
    fi
done
if [[ "$found" == "false" ]]; then
    echo "Tag must be one of: ${allowed_tags[@]}"
    exit 1
fi
if [ "$(docker images -q $image_name:$tag)" ]; then
    docker rmi $image_name:$tag
fi
docker build --no-cache -t $image_name:$tag -f Dockerfile.prod . --platform linux/x86_64
docker tag $image_name:$tag pulzo/$image_name:$tag
docker push pulzo/$image_name:$tag
docker rmi pulzo/$image_name:$tag