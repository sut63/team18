import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';
import SearchIcon from '@material-ui/icons/Search';
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
      to="reserve"
      text="จองห้องพัก"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="fix"
      text="แจ้งซ่อม"
    />
   <SidebarItem
      icon={SearchIcon}
      to="searchdataroom"
      text="ค้นหาห้องพัก"
    />


    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_in"
      text="Sign Out"
      onClick={Clears}
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
