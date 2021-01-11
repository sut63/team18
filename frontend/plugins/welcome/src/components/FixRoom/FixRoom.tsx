import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import {
  MuiPickersUtilsProvider,
  KeyboardTimePicker,
  KeyboardDatePicker,
} from '@material-ui/pickers';
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
import { EntDataRoom } from '../../api/models/EntDataRoom';
import { EntCustomer } from '../../api/models/EntCustomer';
import { EntFurnitureDetail } from '../../api/models/EntFurnitureDetail';
import { Cookies } from '../../Cookie'

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

interface FixRoom {
  Room: number;
  FurnitureDetail: number;
  Customer: number;
  FixDetail: string;
  // create_by: number;
}



const FixRoom: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

   // ดึงคุกกี้
   var ck = new Cookies()
   var cookieName = ck.GetCookie()
   var cookieID = ck.GetID()

   //อันหลักสำหรับสร้าง การ reserve_room
  const [FixRoom, setFixRoom] = React.useState<Partial<FixRoom>>({});

  // data room List
  const [DataRoom, setDataRoom] = React.useState<EntDataRoom[]>([]);
  const getDataRoom = async () => {
    const res = await api.listDataroom({ limit: 10, offset: 0 });
    setDataRoom(res);
  };
  
  

  // furnituredetail List
  const [FurnitureDetail, setFurnitureDetails] = React.useState<EntFurnitureDetail[]>([]);
  const getFurnitureDetail = async () => {
    const res = await api.listFurnituredetail({ limit: 10, offset: 0 });
    setFurnitureDetails(res);
  };

  //customer
  const [customer, setCustomer] = React.useState<EntCustomer>();
  const getCustomer = async () => {
    const res = await api.getCustomer({id: Number(cookieID)});
    setCustomer(res);
  };

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
    getCustomer();
    getDataRoom();
    getFurnitureDetail();
  }, []);

  useEffect(() => {
    setFixRoom({ ...FixRoom, ['Customer']: customer?.id });
  }, [customer]);

  // set data to object reserve_room
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof FixRoom;
    const { value } = event.target;
    setFixRoom({ ...FixRoom, [name]: value });
    console.log(FixRoom);
  };

  // clear input form
  function clear() {
    setFixRoom({});
    
  }

   // clear cookies
  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }
  
  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/fixrooms';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(FixRoom),
    };
   
    console.log(FixRoom); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบแจ้งซ่อม`}>
        <div style={{ marginLeft: 10, marginRight:20 }}>{cookieName}</div>
        <Button
          variant="outlined"
          color="secondary"
          size="large"
          onClick={Clears}
          >
          Logout
        </Button>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled 
                  name="Customer"
                  variant="outlined"
                  value={customer?.name}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>Room</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกห้อง</InputLabel>
                <Select
                  name="Room"
                  value={FixRoom.Room|| ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {DataRoom.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.roomnumber}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>อุปกรณ์ที่ชำรุด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>อุปกรณ์ที่ชำรุด</InputLabel>
                <Select
                  name="FurnitureDetail"
                  value={FixRoom.FurnitureDetail || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {FurnitureDetail.map(item => {
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
              <div className={classes.paper}>รายละเอียดอุปกรณ์ที่ชำรุด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="standard-multiline-static"
                  label="Multiline"
                  multiline
                  rows={4}
                  defaultValue="Default Value"
                  name = " FixDetail "
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการดู
              </Button>
            </Grid>

          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default FixRoom;
