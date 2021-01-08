import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntPromotion } from '../../api/models/EntPromotion'; // import interface User
import { EntTypeRoom } from '../../api/models/EntTypeRoom'; // import interface Resolution
import { EntStatusRoom } from '../../api/models/EntStatusRoom'; // import interface Playlist
import { EntCounterStaff, EntCustomer, EntDataRoom, EntReserveRoom } from '../../api';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface CheckIn {
  Customer: number;
  Counter: number;
  Reserveroom: number;
  Dataroom: number ;
  // create_by: number;
}

const DataRoom: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [CheckIn, setCheckIn] = React.useState<Partial<CheckIn>>({})
  // Customer
  const [customer,setCustomer] = React.useState<EntCustomer[]>([])
  const [dataroom,setDataRoom] = React.useState<EntDataRoom>()
  const getCustomer = async () => {
    const res = await api.listCustomer({})
    setCustomer(res)
  }
  // Counter
  const [ids,setIds] = React.useState<number>(0)

  const [reserveroom,setReserveroom] = React.useState<EntReserveRoom[]>([])
  const getReseveroom = async () => {
    const res = await api.getReserveRoomCustomer({id : ids})
    setReserveroom(res)
  }
  console.log(CheckIn.Reserveroom);
  
  
  const getDataRoom = async () => {
    const res = await api.getDataroom({id : 1})
    setDataRoom(res)
  }

  console.log(reserveroom.map(item => {
    return (
      item.edges?.room?.id
    );
  }));
  


  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  
  // Lifecycle Hooks
  useEffect(() => {
    getCustomer()
  }, []);
  useEffect(() => {
    getDataRoom()
   
  }, [CheckIn.Reserveroom]);
  useEffect(() => {

    getReseveroom()
  }, [CheckIn.Customer]);

  // set data to object Checkin
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof CheckIn;
    const { value } = event.target;
    setCheckIn({ ...CheckIn, [name]: value });
    setIds(event.target.value)
    setDataRoom(event.target.value)

  };
  console.log(reserveroom); 

  // clear input form
 

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/checkins ';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(DataRoom),
    };

    

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.id != null) {
          //clear();
          Toast.fire({
            icon: 'success',
            title: 'Ckeck in สำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'error',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบย่อย Check In`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Tanapon Kongjaroensuk</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
                      
            <Grid item xs={3}>
              <div className={classes.paper}>customer</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>select customer</InputLabel>
                <Select
                  name="Customer"
                  value={CheckIn.Customer || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {customer.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>reserveroom</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>select reserveroom</InputLabel>
                <Select
                  name="Reserveroom"
                  value={CheckIn.Reserveroom || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {reserveroom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Room Number</div>
            </Grid>
            <Grid item xs={9}>
            <TextField  
            name = "Price"
            variant="outlined" 
            value ={dataroom?.roomnumber}
            />
            </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}>ราคาห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
            <TextField  
            label="price" 
            name = "Price"
            variant="outlined" 
            onChange={handleChange}/>
           </Grid>
           <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                onClick={save}
              >
                Check in
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default DataRoom;
