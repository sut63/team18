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
    EntDataRoomEdges,
    EntDataRoomEdgesFromJSON,
    EntDataRoomEdgesFromJSONTyped,
    EntDataRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDataRoom
 */
export interface EntDataRoom {
    /**
     * 
     * @type {EntDataRoomEdges}
     * @memberof EntDataRoom
     */
    edges?: EntDataRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDataRoom
     */
    id?: number;
    /**
     * Price holds the value of the "price" field.
     * @type {number}
     * @memberof EntDataRoom
     */
    price?: number;
    /**
     * Roomnumber holds the value of the "roomnumber" field.
     * @type {string}
     * @memberof EntDataRoom
     */
    roomnumber?: string;
}

export function EntDataRoomFromJSON(json: any): EntDataRoom {
    return EntDataRoomFromJSONTyped(json, false);
}

export function EntDataRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDataRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntDataRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'price': !exists(json, 'price') ? undefined : json['price'],
        'roomnumber': !exists(json, 'roomnumber') ? undefined : json['roomnumber'],
    };
}

export function EntDataRoomToJSON(value?: EntDataRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntDataRoomEdgesToJSON(value.edges),
        'id': value.id,
        'price': value.price,
        'roomnumber': value.roomnumber,
    };
}

