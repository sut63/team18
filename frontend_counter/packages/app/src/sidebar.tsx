import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

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

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
