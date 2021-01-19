/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntReserveRoomEdges,
    EntReserveRoomEdgesFromJSON,
    EntReserveRoomEdgesFromJSONTyped,
    EntReserveRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntReserveRoom
 */
export interface EntReserveRoom {
    /**
     * Amount holds the value of the "amount" field.
     * @type {number}
     * @memberof EntReserveRoom
     */
    amount?: number;
    /**
     * Duration holds the value of the "duration" field.
     * @type {number}
     * @memberof EntReserveRoom
     */
    duration?: number;
    /**
     * 
     * @type {EntReserveRoomEdges}
     * @memberof EntReserveRoom
     */
    edges?: EntReserveRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntReserveRoom
     */
    id?: number;
    /**
     * NetPrice holds the value of the "net_price" field.
     * @type {number}
     * @memberof EntReserveRoom
     */
    netPrice?: number;
    /**
     * Province holds the value of the "province" field.
     * @type {string}
     * @memberof EntReserveRoom
     */
    province?: string;
    /**
     * ReserveDate holds the value of the "reserve_date" field.
     * @type {string}
     * @memberof EntReserveRoom
     */
    reserveDate?: string;
    /**
     * Tel holds the value of the "tel" field.
     * @type {string}
     * @memberof EntReserveRoom
     */
    tel?: string;
}

export function EntReserveRoomFromJSON(json: any): EntReserveRoom {
    return EntReserveRoomFromJSONTyped(json, false);
}

export function EntReserveRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntReserveRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'duration': !exists(json, 'duration') ? undefined : json['duration'],
        'edges': !exists(json, 'edges') ? undefined : EntReserveRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'netPrice': !exists(json, 'net_price') ? undefined : json['net_price'],
        'province': !exists(json, 'province') ? undefined : json['province'],
        'reserveDate': !exists(json, 'reserve_date') ? undefined : json['reserve_date'],
        'tel': !exists(json, 'tel') ? undefined : json['tel'],
    };
}

export function EntReserveRoomToJSON(value?: EntReserveRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'amount': value.amount,
        'duration': value.duration,
        'edges': EntReserveRoomEdgesToJSON(value.edges),
        'id': value.id,
        'net_price': value.netPrice,
        'province': value.province,
        'reserve_date': value.reserveDate,
        'tel': value.tel,
    };
}


