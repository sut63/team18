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
    EntTypeRoomEdges,
    EntTypeRoomEdgesFromJSON,
    EntTypeRoomEdgesFromJSONTyped,
    EntTypeRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTypeRoom
 */
export interface EntTypeRoom {
    /**
     * 
     * @type {EntTypeRoomEdges}
     * @memberof EntTypeRoom
     */
    edges?: EntTypeRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTypeRoom
     */
    id?: number;
    /**
     * TypeName holds the value of the "type_name" field.
     * @type {string}
     * @memberof EntTypeRoom
     */
    typeName?: string;
}

export function EntTypeRoomFromJSON(json: any): EntTypeRoom {
    return EntTypeRoomFromJSONTyped(json, false);
}

export function EntTypeRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypeRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntTypeRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'typeName': !exists(json, 'type_name') ? undefined : json['type_name'],
    };
}

export function EntTypeRoomToJSON(value?: EntTypeRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntTypeRoomEdgesToJSON(value.edges),
        'id': value.id,
        'type_name': value.typeName,
    };
}


