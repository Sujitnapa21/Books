/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
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
    EntBooks,
    EntBooksFromJSON,
    EntBooksFromJSONTyped,
    EntBooksToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomamountEdges
 */
export interface EntRoomamountEdges {
    /**
     * Books holds the value of the books edge.
     * @type {Array<EntBooks>}
     * @memberof EntRoomamountEdges
     */
    books?: Array<EntBooks>;
}

export function EntRoomamountEdgesFromJSON(json: any): EntRoomamountEdges {
    return EntRoomamountEdgesFromJSONTyped(json, false);
}

export function EntRoomamountEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomamountEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'books': !exists(json, 'books') ? undefined : ((json['books'] as Array<any>).map(EntBooksFromJSON)),
    };
}

export function EntRoomamountEdgesToJSON(value?: EntRoomamountEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'books': value.books === undefined ? undefined : ((value.books as Array<any>).map(EntBooksToJSON)),
    };
}


