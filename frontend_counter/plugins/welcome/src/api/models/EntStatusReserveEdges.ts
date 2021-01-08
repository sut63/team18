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
    EntReserveRoom,
    EntReserveRoomFromJSON,
    EntReserveRoomFromJSONTyped,
    EntReserveRoomToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusReserveEdges
 */
export interface EntStatusReserveEdges {
    /**
     * Reserves holds the value of the reserves edge.
     * @type {Array<EntReserveRoom>}
     * @memberof EntStatusReserveEdges
     */
    reserves?: Array<EntReserveRoom>;
}

export function EntStatusReserveEdgesFromJSON(json: any): EntStatusReserveEdges {
    return EntStatusReserveEdgesFromJSONTyped(json, false);
}

export function EntStatusReserveEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusReserveEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reserves': !exists(json, 'reserves') ? undefined : ((json['reserves'] as Array<any>).map(EntReserveRoomFromJSON)),
    };
}

export function EntStatusReserveEdgesToJSON(value?: EntStatusReserveEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reserves': value.reserves === undefined ? undefined : ((value.reserves as Array<any>).map(EntReserveRoomToJSON)),
    };
}


