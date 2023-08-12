'use client'

import React from 'react';
import Drawer from '@mui/material/Drawer';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import GroupsIcon from '@mui/icons-material/Groups';
import DnsIcon from '@mui/icons-material/Dns';

type menuListProps = {
  onSelected:(key:string) => void
}

const MenuList = (props:menuListProps) => {
  return (
  <Box
    role="presentation"
  >
    <List>
    <ListItem disablePadding>
      <ListItemButton onClick={()=>{props.onSelected("matches")}}> 
        <ListItemIcon>
          <GroupsIcon />
        </ListItemIcon>
        <ListItemText primary="MATCHES" />
      </ListItemButton>
    </ListItem>
    <ListItem disablePadding>
    <ListItemButton onClick={()=>{props.onSelected("servers")}}> 
        <ListItemIcon>
          <DnsIcon />
        </ListItemIcon>
        <ListItemText primary="SERVERS" />
      </ListItemButton>
    </ListItem>
    </List>
  </Box>
  )
};

function MenuDrawer() {
  const [selected, setSelected] = React.useState<string>("matches")
  return (
    <Drawer open={true}>
      <MenuList onSelected={setSelected} />
      {selected}
    </Drawer>
  )
}

export default function Home() {
  return (
    <MenuDrawer />
  )
}
