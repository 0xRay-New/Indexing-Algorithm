# Indexing-Algorithm

This repository contains all of the code used by 0xRay to index NFTs.

**Indexing Algorithm**

![alt text](https://media.discordapp.net/attachments/893372833743372321/1096589034047942756/image.png?width=136&height=62)

**Variables**
- `v`: the total number of traits in a given NFT collection.
- `N`: the total number of traits, inclusive of a nonetrait*. 
- `T`: the total number of NFTs in the collection.
- `B`: the number of traits for a given individual NFT in the collection.
- `a`: the number of NFTs with each trait in the collection.

**Clarifications**
- Not all NFTs in a given collection will have a trait. For example, in a collection of 10 NFTs, 5 may have a blue shirt, 4 may have a red shirt, and 1 may not have a shirt at all. The collection metadata may not expressly state that 1 NFT has no shirt, and only mentions the 5 blue shirts and 4 red shirts. The nonetrait simply makes this fact expressly stated. 

**Important functions**
- [Get collectionwide metadata (including generating nonetraits)](calculator/getAssetsRequest.go)
- [Get metadata for each NFT](calculator/getMetadata.go)

**Potential future improvements**
- Speed up the indexing process by having multiple goroutines at a time access the OpenSea metadata API.

*A special thanks to [Ishaan](https://twitter.com/ishaanganti) for helping develop the algorithm.*