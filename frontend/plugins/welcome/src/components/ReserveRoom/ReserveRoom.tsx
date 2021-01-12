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
import { EntPromotion } from '../../api/models/EntPromotion';
import { EntCustomer } from '../../api/models/EntCustomer';
import { EntStatusReserve } from '../../api/models/EntStatusReserve';
import { Cookies } from '../../Cookie'
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';

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

interface reserve {
  Rooms: number;
  Promotions: number;
  Customers: number;
  Status: number;
  ReserveDate: Date;
  OutDate: Date;
  NetPrice: number;
  // create_by: number;
}

interface room {
  StatusRoom: number;
  // create_by: number;
}

const ReserveRoom: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  // ดึงคุกกี้
  var ck = new Cookies()
  var cookieName = ck.GetCookie()
  var cookieID = ck.GetID()

  //อันหลักสำหรับสร้าง การ reserve_room
  const [reserve_room, setReserveroom] = React.useState<Partial<reserve>>({});

  // data room List
  const [dataroom, setdataroom] = React.useState<EntDataRoom[]>([]);
  const getDataRoom = async () => {
    const res = await api.listDataRoomPromo({ id: ids });
    setdataroom(res);
  };
  // Room
  const [room, setroom] = React.useState<EntDataRoom>();
  const getRoom = async () => {
    const res = await api.getDataroom({ id: dids })
    setroom(res);
    setPrice(res.price);
  }

  // promotion List
  const [promotions, setPromotions] = React.useState<EntPromotion[]>([]);
  const getPromotion = async () => {
    const res = await api.listPromotion({ limit: 10, offset: 0 });
    setPromotions(res);
  };
  // promotion
  const [promo, setPromo] = React.useState<EntPromotion>();
  const getPromo = async () => {
    const res = await api.getPromotion({ id: ids })
    setPromo(res);
    setDiscount(res.discount);
    setPrice(0);
    clearNet();
  }

  //customer
  const [customers, setCustomers] = React.useState<EntCustomer>();
  const getCustomes = async () => {
    const res = await api.getCustomer({ id: Number(cookieID) });
    setCustomers(res);
  };

  //status
  const [status, setStatus] = React.useState<EntStatusReserve>();
  const getStatus = async () => {
    const res = await api.getStatusReserve({ id: 1 })
    setStatus(res);
  }

  //find net price
  const [price, setPrice] = React.useState<any>(0)
  const [discount, setDiscount] = React.useState<any>(0)
  const [netprice, setNetprice] = React.useState<any>(0)
  const getNetprice = async () => {
    setNetprice(price - discount);
  };

  //set time
  const [timeIn, setTimeIn] = React.useState<any>(0)
  const [timeOut, setTimeOut] = React.useState<any>(0)

  //id for find room and price
  const [ids, setIds] = React.useState<number>(0)
  const [dids, setDIds] = React.useState<number>(0)

  // alert setting
  const [open, setOpen] = React.useState(false);
  const [fail, setFail] = React.useState(false);

  // Lifecycle Hooks
  useEffect(() => {
    getCustomes();
    getPromotion();
    getStatus();
  }, []);

  useEffect(() => {
    getPromo();
    getDataRoom();
  }, [reserve_room.Promotions]);

  useEffect(() => {
    getRoom();
  }, [reserve_room.Rooms]);

  useEffect(() => {
    getNetprice();
  }, [price]);

  useEffect(() => {
    setReserveroom({ ...reserve_room, ['NetPrice']: netprice });
  }, [netprice]);

  useEffect(() => {
    setReserveroom({ ...reserve_room, ['Status']: status?.id });
  }, [status]);

  useEffect(() => {
    setReserveroom({ ...reserve_room, ['Customers']: customers?.id });
  }, [customers]);

  // set data to object reserve_room
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof ReserveRoom;
    const { value } = event.target;
    setReserveroom({ ...reserve_room, [name]: value });
    setIds(event.target.value);
    setDIds(event.target.value);
    console.log(reserve_room);
  };

  // set data for time in
  const handleChangeTimeIN = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof ReserveRoom;
    const { value } = event.target;
    const time = value + ":00+07:00"
    setReserveroom({ ...reserve_room, [name]: time });
    setTimeIn(value);
    console.log(reserve_room);
  };

  // set data for time out
  const handleChangeTimeOut = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof ReserveRoom;
    const { value } = event.target;
    const time = value + ":00+07:00"
    setReserveroom({ ...reserve_room, [name]: time });
    setTimeOut(value);
    console.log(reserve_room);
  };

  //close alert 
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === 'clickaway') {
      return;
    }
    setFail(false);
    setOpen(false);
  };

  // clear input form
  function clear() {
    setReserveroom({});
    setDIds(0);
    setIds(0);
    setDiscount(0);
    setPrice(0);
    setNetprice(0);
    setTimeIn({});
    setTimeOut({});
    getCustomes();
  }

  // clear netprice form
  function clearNet() {
    setNetprice(0);
  }

  // clear cookies
  function Clears() {
    ck.ClearCookie()
    window.location.reload(false)
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/ReserveRooms';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(reserve_room),
    };

    console.log(reserve_room); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.status);
        if (data.status === true) {
          clear();
          setOpen(true);
        } else {
          setFail(true);
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบจองห้องพัก`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
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
                  value={customers?.name}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>promotion</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกโปรโมชั่น</InputLabel>
                <Select
                  name="Promotions"
                  value={reserve_room.Promotions || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {promotions.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.promotionName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>room</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกห้องพัก</InputLabel>
                <Select
                  name="Rooms"
                  value={reserve_room.Rooms || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {dataroom.map(item => {
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
              <div className={classes.paper}>เลือกวันเข้าพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="ReserveDate"
                  type="datetime-local"
                  value={timeIn}
                  defaultValue="2020-12-31"
                  onChange={handleChangeTimeIN}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกวันออก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  name="OutDate"
                  type="datetime-local"
                  value={timeOut}
                  defaultValue="2020-12-31"
                  onChange={handleChangeTimeOut}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ราคา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Price"
                  variant="outlined"
                  value={price}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ส่วนลด</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Discount"
                  variant="outlined"
                  value={discount}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ราคาสุทธิ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="NetPrice"
                  variant="outlined"
                  value={netprice}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>สถานะ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  disabled
                  name="Status"
                  variant="outlined"
                  value={status?.statusName}
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

          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              This is a success message!
        </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              This is a error message!
        </Alert>
          </Snackbar>

        </Container>
      </Content>
    </Page>
  );
};

export default ReserveRoom;
