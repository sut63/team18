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

const ReserveRoom: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [reserve_room, setReserveroom] = React.useState<Partial<reserve>>({});

  const [dataroom, setdataroom] = React.useState<EntDataRoom[]>([]);
  const [room, setroom] = React.useState<EntDataRoom>();

  const [promotions, setPromotions] = React.useState<EntPromotion[]>([]);
  const [promo, setPromo] = React.useState<EntPromotion>();

  const [customers, setCustomers] = React.useState<EntCustomer[]>([]);

  const [status, setStatus] = React.useState<EntStatusReserve>();

  const [price, setPrice] = React.useState<any>(0)
  const [discount, setDiscount] = React.useState<any>(0)
  const [netprice, setNetprice] = React.useState<any>(0)

  const [timeIn, setTimeIn] = React.useState<any>(0)
  const [timeOut, setTimeOut] = React.useState<any>(0)

  const [ids, setIds] = React.useState<number>(0)
  const [dids, setDIds] = React.useState<number>(0)

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

  const getCustomes = async () => {
    const res = await api.listCustomer({ limit: 10, offset: 0 });
    setCustomers(res);
  };

  const getDataRoom = async () => {
    const res = await api.listDataRoomPromo({ id: ids });
    setdataroom(res);
  };

  const getPromotion = async () => {
    const res = await api.listPromotion({ limit: 10, offset: 0 });
    setPromotions(res);
  };

  const getRoom = async () => {
    const res = await api.getDataroom({ id: dids })
    setroom(res);
    setPrice(res.price);
  }

  const getStatus = async () => {
    const res = await api.getStatusReserve({ id: 1 })
    setStatus(res);
  }

  const getPromo = async () => {
    const res = await api.getPromotion({ id: ids })
    setPromo(res);
    setDiscount(res.discount);
    setPrice(0);
    clearNet();
  }

  const getNetprice = async () => {
    setNetprice(price - discount);
  };

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


  // set data to object playlist_video
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

  // clear input form
  function clear() {
    setReserveroom({});
  }

  // clear netprice form
  function clearNet() {
    setNetprice(0);
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
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>รอทำ login</div>
      </Header>
    </Page>
  );
};

export default ReserveRoom;
