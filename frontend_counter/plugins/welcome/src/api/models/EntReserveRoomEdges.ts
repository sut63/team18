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
    EntCheckIn,
    EntCheckInFromJSON,
    EntCheckInFromJSONTyped,
    EntCheckInToJSON,
    EntCustomer,
    EntCustomerFromJSON,
    EntCustomerFromJSONTyped,
    EntCustomerToJSON,
    EntDataRoom,
    EntDataRoomFromJSON,
    EntDataRoomFromJSONTyped,
    EntDataRoomToJSON,
    EntPromotion,
    EntPromotionFromJSON,
    EntPromotionFromJSONTyped,
    EntPromotionToJSON,
    EntStatusReserve,
    EntStatusReserveFromJSON,
    EntStatusReserveFromJSONTyped,
    EntStatusReserveToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntReserveRoomEdges
 */
export interface EntReserveRoomEdges {
    /**
     * Checkins holds the value of the checkins edge.
     * @type {Array<EntCheckIn>}
     * @memberof EntReserveRoomEdges
     */
    checkins?: Array<EntCheckIn>;
    /**
     * 
     * @type {EntCustomer}
     * @memberof EntReserveRoomEdges
     */
    customer?: EntCustomer;
    /**
     * 
     * @type {EntPromotion}
     * @memberof EntReserveRoomEdges
     */
    promotion?: EntPromotion;
    /**
     * 
     * @type {EntDataRoom}
     * @memberof EntReserveRoomEdges
     */
    room?: EntDataRoom;
    /**
     * 
     * @type {EntStatusReserve}
     * @memberof EntReserveRoomEdges
     */
    status?: EntStatusReserve;
}

export function EntReserveRoomEdgesFromJSON(json: any): EntReserveRoomEdges {
    return EntReserveRoomEdgesFromJSONTyped(json, false);
}

export function EntReserveRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntReserveRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'checkins': !exists(json, 'Checkins') ? undefined : ((json['Checkins'] as Array<any>).map(EntCheckInFromJSON)),
        'customer': !exists(json, 'Customer') ? undefined : EntCustomerFromJSON(json['Customer']),
        'promotion': !exists(json, 'Promotion') ? undefined : EntPromotionFromJSON(json['Promotion']),
        'room': !exists(json, 'Room') ? undefined : EntDataRoomFromJSON(json['Room']),
        'status': !exists(json, 'Status') ? undefined : EntStatusReserveFromJSON(json['Status']),
    };
}

export function EntReserveRoomEdgesToJSON(value?: EntReserveRoomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'checkins': value.checkins === undefined ? undefined : ((value.checkins as Array<any>).map(EntCheckInToJSON)),
        'customer': EntCustomerToJSON(value.customer),
        'promotion': EntPromotionToJSON(value.promotion),
        'room': EntDataRoomToJSON(value.room),
        'status': EntStatusReserveToJSON(value.status),
    };
}


