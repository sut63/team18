import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import SearchIcon from '@material-ui/icons/Search';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';
import { Cookies } from 'plugin-welcome/src/Cookie'

import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

var ck = new Cookies()
 // clear cookies
 function Clears() {
  ck.ClearCookie()
  window.location.reload(false)
}

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem
      icon={CreateComponentIcon}
      to="dataroom"
      text="Data Room"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="CheckIn"
      text="Check In"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="checkout"
      text="Check Out"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="FurnitureDetail"
      text="FurnitureDetail"
    />
    <SidebarItem
      icon={SearchIcon}
      to="SearchReserveRoom"
      text="ค้นหาใบจอง"
    />
    <SidebarItem
      icon={SearchIcon}
      to="SearchCheckIn"
      text="ค้นหาการ check in"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to=" "
      text="Sign Out"
      onClick={Clears}
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
