import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntDataRoom } from '../../api/models/';
const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
});

export default function TablesRoom() {
    const classes = useStyles();
    const api = new DefaultApi();

    const [rooms, setRooms] = useState<EntDataRoom[]>(Array);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getRoom = async () => {
            const res = await api.listDataroom({ limit: 20, offset: 0 });
            setLoading(false);
            setRooms(res);
        };
        getRoom();
    }, [loading]);

    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="center">หมายเลขห้อง</TableCell>
                        <TableCell align="center">ประเภทห้อง</TableCell>
                        <TableCell align="center">สถานะ</TableCell>
                        <TableCell align="center">ราคา</TableCell>
                        <TableCell align="center">โปรโมชั่น</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rooms.map((item: EntDataRoom) => (
                        <TableRow key={item.id}>
                            <TableCell align="center">{item.roomnumber}</TableCell>
                            <TableCell align="center">{item.edges?.typeroom?.typeName}</TableCell>
                            <TableCell align="center">{item.edges?.statusroom?.statusName}</TableCell>
                            <TableCell align="center">{item.price} บาท</TableCell>
                            <TableCell align="center">{item.edges?.promotion?.promotionName}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}