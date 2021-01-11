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
    EntStatusReserveEdges,
    EntStatusReserveEdgesFromJSON,
    EntStatusReserveEdgesFromJSONTyped,
    EntStatusReserveEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStatusReserve
 */
export interface EntStatusReserve {
    /**
     * 
     * @type {EntStatusReserveEdges}
     * @memberof EntStatusReserve
     */
    edges?: EntStatusReserveEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStatusReserve
     */
    id?: number;
    /**
     * StatusName holds the value of the "status_name" field.
     * @type {string}
     * @memberof EntStatusReserve
     */
    statusName?: string;
}

export function EntStatusReserveFromJSON(json: any): EntStatusReserve {
    return EntStatusReserveFromJSONTyped(json, false);
}

export function EntStatusReserveFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStatusReserve {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntStatusReserveEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'statusName': !exists(json, 'status_name') ? undefined : json['status_name'],
    };
}

export function EntStatusReserveToJSON(value?: EntStatusReserve | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntStatusReserveEdgesToJSON(value.edges),
        'id': value.id,
        'status_name': value.statusName,
    };
}


