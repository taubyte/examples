import { ethers } from 'ethers';
import MintABI from '../mint_abi.json';
import ComputeABI from '../compute_abi.json';

let provider;
if (window.ethereum) {
    provider = new ethers.providers.Web3Provider(window.ethereum);
} else {
    console.log('Please install MetaMask!');
}
const signer = provider.getSigner();

const MintContract = new ethers.Contract('0x017E5f589795ffc218E97D2B82ECf4c7bd387e63', MintABI, signer);
const ComputeContract = new ethers.Contract('0xE91dCD50913066Db9E7Ea8D9F6D41138163CE90a', ComputeABI, signer);

export { provider, MintContract, ComputeContract };
