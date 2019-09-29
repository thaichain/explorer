import {MenuItem} from '../models/menu_item.model';
import {ErcName, InterfaceName, ThemeColor} from './enums';

export const THEME_SETTINGS = {
  color: ThemeColor.LIGHT,
};

export const LOGO_NAMES = {
  [ThemeColor.LIGHT]: 'logo_fullcolor.svg',
  [ThemeColor.DARK]: 'logo_allwhite.svg',
};

export const ROUTES = {
  HOME: 'home',
  BLOCK: 'block',
  ADDRESS_FULL: 'address',
  ADDRESS: 'addr',
  TOKEN: 'token',
  RICHLIST: 'richlist',
  TRANSACTION: 'tx',
  SETTINGS: 'settings',
  VERIFY: 'verify',
  WALLET: 'wallet',
  CREATE_WALLET: 'create-account',
  SEND_TX: 'send-tx',
  SIGNERS: 'signers',
  CONTRACTS: 'contracts',
};

export const MENU_ITEMS: MenuItem[] = [
  {
    title: 'Main',
    link: ROUTES.HOME,
  },
 /*
  {
    title: 'Rich List',
    link: ROUTES.RICHLIST,
  },
 */
  {
    title: 'Verify Contract',
    link: '/verify',
  },
  {
    title: 'Contracts',
    link: ROUTES.CONTRACTS,
  },
  {
    title: 'Wallet',
    link: ROUTES.WALLET,
  },
  {
    title: 'Signers',
    link: ROUTES.SIGNERS,
  },
 /* 
  {
    title: 'Network Stats',
    link: 'https://stats.tch.in.th/',
    external: true
  },
  {
    title: 'Settings',
    link: '/settings',
  },*/
];

export const DEFAULT_GAS_LIMIT = 136500000;

export const TOKEN_TYPES = {
  Go20: 'ERC20',
  Go20Burnable: 'ERC20 Burnable',
  Go20Capped: 'ERC20 Capped',
  Go20Detailed: 'ERC20 Detailed',
  Go20Mintable: 'ERC20 Mintable',
  Go20Pausable: 'ERC20 Pausable',
  Go165: 'ERC165',
  Go721: 'ERC721',
  Go721Burnable: 'ERC721 Burnable',
  Go721Receiver: 'ERC721 Receiver',
  Go721Metadata: 'ERC721 Metadata',
  Go721Enumerable: 'ERC721 Enumerable',
  Go721Pausable: 'ERC721 Pausable',
  Go721Mintable: 'ERC721 Mintable',
  Go721MetadataMintable: 'ERC721 Metadata Mintable',
  Go721Full: 'ERC721 Full',
  Go820: 'ERC820',
  Go1155: 'ERC1155',
  Go1155Receiver: 'ERC1155 Receiver',
  Go1155Metadata: 'ERC1155 Metadata',
  Go223: 'ERC223',
  Go223Receiver: 'ERC223 Receiver',
  Go621: 'ERC621',
  Go777: 'ERC777',
  Go777Receiver: 'ERC777 Receiver',
  Go777Sender: 'ERC777 Sender',
  Go827: 'ERC827',
  Go884: 'ERC884',
};

export const ERC_INTERFACE_IDENTIFIERS = {
  [ErcName.Go20]: [InterfaceName.Allowance, InterfaceName.Approve, InterfaceName.BalanceOf, InterfaceName.TotalSupply, InterfaceName.Transfer, InterfaceName.TransferFrom],
  [ErcName.Go721]: [InterfaceName.Approve, InterfaceName.BalanceOf, InterfaceName.GetApproved, InterfaceName.IsApprovedForAll, InterfaceName.OwnerOf, InterfaceName.SafeTransferFrom, InterfaceName.SafeTransferFrom1, InterfaceName.SetApprovalForAll, InterfaceName.TransferFrom],
};

export const TOKEN_ABI_NAMES: string[] = ['totalSupply', 'balanceOf'];

export const META_TITLES = {
  DEFAULT: {
    title: 'ThaiChain Explorer',
  },
  HOME: {
    title: 'Home',
  },
  BLOCK: {
    title: 'Block',
  },
  ADDRESS: {
    title: 'Address',
  },
  CONTRACT: {
    title: 'Contract',
  },
  TOKEN: {
    title: 'Token',
  },
  RICHLISLT: {
    title: 'Richlist',
  },
  CONTRACTS: {
    title: 'Contracts List',
  },
  TRANSACTION: {
    title: 'Transaction',
  },
  VERIFY: {
    title: 'Verify contract',
  },
  SEND_TX: {
    title: 'Send transaction',
  },
  SIGNERS: {
    title: 'Signers',
  },
  WALLET: {
    title: 'Wallet',
  },
  CREATE_WALLET: {
    title: 'Create account',
  },
  SEND_WALLET: {
    title: 'Send GO',
  },
  DEPLOY_CONTRACT: {
    title: 'Deploy contract',
  },
  USE_CONTRACT: {
    title: 'Interact with a Smart Contract',
  },
  OPEN_WALLET: {
    title: 'Open wallet',
  },
  NOT_FOUND: {
    title: 'Not found',
  },
};
