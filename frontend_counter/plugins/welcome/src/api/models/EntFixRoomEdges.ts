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
    EntCustomer,
    EntCustomerFromJSON,
    EntCustomerFromJSONTyped,
    EntCustomerToJSON,
    EntDataRoom,
    EntDataRoomFromJSON,
    EntDataRoomFromJSONTyped,
    EntDataRoomToJSON,
    EntFurnitureDetail,
    EntFurnitureDetailFromJSON,
    EntFurnitureDetailFromJSONTyped,
    EntFurnitureDetailToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFixRoomEdges
 */
export interface EntFixRoomEdges {
    /**
     * 
     * @type {EntCustomer}
     * @memberof EntFixRoomEdges
     */
    customer?: EntCustomer;
    /**
     * 
     * @type {EntFurnitureDetail}
     * @memberof EntFixRoomEdges
     */
    furnitureDetail?: EntFurnitureDetail;
    /**
     * 
     * @type {EntDataRoom}
     * @memberof EntFixRoomEdges
     */
    room?: EntDataRoom;
}

export function EntFixRoomEdgesFromJSON(json: any): EntFixRoomEdges {
    return EntFixRoomEdgesFromJSONTyped(json, false);
}

export function EntFixRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFixRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'customer': !exists(json, 'Customer') ? undefined : EntCustomerFromJSON(json['Customer']),
        'furnitureDetail': !exists(json, 'FurnitureDetail') ? undefined : EntFurnitureDetailFromJSON(json['FurnitureDetail']),
        'room': !exists(json, 'Room') ? undefined : EntDataRoomFromJSON(json['Room']),
    };
}

export function EntFixRoomEdgesToJSON(value?: EntFixRoomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Customer': EntCustomerToJSON(value.customer),
        'FurnitureDetail': EntFurnitureDetailToJSON(value.furnitureDetail),
        'Room': EntDataRoomToJSON(value.room),
    };
}


