import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
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
import { EntCustomer, EntReserveRoom } from '../../api';
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
    bottom: {
        marginLeft: theme.spacing(3),
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(2),
    },
    divider: {
        margin: theme.spacing(2, 0),
    },
    table: {
        minWidth: 650,
    },
}));


const SearchReserveRoom: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    // ดึงคุกกี้
    var ck = new Cookies()
    var cookieName = ck.GetCookie()
    var cookieID = ck.GetID()

    // Customer
    const [idCustomer, setIdcustomer] = React.useState<number>(0)
    const [customer, setCustomer] = React.useState<EntCustomer[]>([])
    const getCustomer = async () => {
        const res = await api.listCustomer({})
        setCustomer(res)
    }

    // ReserveRoom  
    const [reserveroom, setReserveroom] = React.useState<EntReserveRoom[]>([])
    const getReseveroom = async () => {
        const res = await api.getReserveRoomCustomer({ id: idCustomer })
        setReserveroom(res)
    }
    // Lifecycle Hooks
    useEffect(() => {
        getCustomer();
    }, []);

    // // Lifecycle Hooks
    // useEffect(() => {
    //     getReseveroom();
    // }, [idCustomer]);

    // set data to object furniture_detail
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const name = event.target.name as keyof typeof SearchReserveRoom;
        const { value } = event.target;
        setIdcustomer(value);
    };

    // clear cookies
    function Clears() {
        ck.ClearCookie()
        window.location.reload(false)
    }

    // function seach data
    function seach() {
        getReseveroom();
    }

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหาใบจอง`}>
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
                <Grid container spacing={1}>
                    <Grid item xs={1}>
                        <div className={classes.paper}>customer</div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>select customer</InputLabel>
                            <Select
                                name="Customer"
                                value={idCustomer || ''} // (undefined || '') = ''
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
                        <Button
                            className={classes.bottom}
                            variant="contained"
                            color="primary"
                            size="large"
                            onClick={seach}
                        >
                            ค้นหา
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตารางใบจอง:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">เลขใบจอง</TableCell>
                                <TableCell align="center">ห้องที่จอง</TableCell>
                                <TableCell align="center">สถานะ</TableCell>
                                <TableCell align="center">จำนวนคนที่เข้าพัก</TableCell>
                                <TableCell align="center">เบอร์โทรติดต่อ</TableCell>
                                <TableCell align="center">วันที่เข้าพัก</TableCell>
                                <TableCell align="center">ราคาสุทธิ</TableCell>
                                <TableCell align="center">request</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {reserveroom.map((item: EntReserveRoom) => (
                                <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.room?.roomnumber}</TableCell>
                                    <TableCell align="center">{item.edges?.status?.statusName}</TableCell>
                                    <TableCell align="center">{item.amount}</TableCell>
                                    <TableCell align="center">{item.phoneNumber}</TableCell>
                                    <TableCell align="center">{item.reserveDate}</TableCell>
                                    <TableCell align="center">{item.netPrice} บาท</TableCell>
                                    <TableCell align="center">{item.request}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Content>
        </Page>
    );
};

export default SearchReserveRoom;
