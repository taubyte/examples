// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol"; 

contract TauFRC721 is ERC721URIStorage {
/** @notice Counter keeps track of the token ID number for each unique NFT minted in the NFT collection */
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

/** @notice This struct stores information about each NFT minted */
    struct tauFRC721NFT {
        address owner;
        string tokenURI;
        uint256 tokenId;
    }

 /** @notice Keeping an array for each of the NFT's minted on this contract allows me to get information on them all with a read-only front end call */
    tauFRC721NFT[] public nftCollection;
/** @notice The mapping allows me to find NFT's owned by a particular wallet address. I'm only handling the case where an NFT is minted to an owner in this contract - but you'd need to handle others in a mainnet contract like sending to other wallets */
    mapping(address => tauFRC721NFT[]) public nftCollectionByOwner;

/** @notice This event will be triggered (emitted) each time a new NFT is minted - which I will watch for on my front end in order to load new information that comes in about the collection as it happens */
    event NewTauFRC721NFTMinted(
      address  sender,
      uint256  tokenId,
      string tokenURI
    );

/** @notice Creates the NFT Collection Contract with a Name and Symbol */
    constructor() ERC721("Taubyte NFTs", "TAU") {

    }

/** 
@notice The main function which will mint each NFT.
The ipfsURI is a link to the ipfs content identifier hash of the NFT metadata stored on NFT.Storage. This data minimally includes name, description and the image in a JSON.
*/
    function mintTauNFT(address owner, string memory ipfsURI)
        public
        returns (uint256)
    {
        // get the tokenID for this new NFT
        uint256 newItemId = _tokenIds.current();

        // Format info for saving to our array
        tauFRC721NFT memory newNFT = tauFRC721NFT({
            owner: msg.sender,
            tokenURI: ipfsURI,
            tokenId: newItemId
        });

        //mint the NFT to the chain
        _mint(owner, newItemId);
        //Set the NFT Metadata for this NFT
        _setTokenURI(newItemId, ipfsURI);

        _tokenIds.increment();

        //Add it to our collection array & owner mapping
        nftCollection.push(newNFT);
        nftCollectionByOwner[owner].push(newNFT);

        // Emit an event on-chain to say we've minted an NFT
        emit NewTauFRC721NFTMinted(
          msg.sender,
          newItemId,
          ipfsURI
        );

        return newItemId;
    }

    /**
     * @notice helper function to display NFTs for frontends
     */
    function getNFTCollection() public view returns (tauFRC721NFT[] memory) {
        return nftCollection;
    }

    /**
     * @notice helper function to fetch NFT's by owner
     */
    function getNFTCollectionByOwner(address owner) public view returns (tauFRC721NFT[] memory){
        return nftCollectionByOwner[owner];
    }
}