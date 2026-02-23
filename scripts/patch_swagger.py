#!/usr/bin/env python3
import json
import sys


LIST_INVITATIONS_PATH = "/user-invitations"
DELETE_INVITATION_PATH = "/user-invitations/{email}"
SEND_INVITATION_PATH = "/organizations/{id}/invite"


def ensure_object_schema(components, name):
    schemas = components.setdefault("schemas", {})
    if name in schemas:
        return
    schemas[name] = {
        "type": "object",
        "properties": {
            "uuid": {"type": "string"},
            "code": {"type": "integer"},
            "message": {"type": "string"},
            "name": {"type": "string"},
            "httpStatusCode": {"type": "integer"},
        },
        "required": ["code", "message"],
    }


def ensure_invitation_schemas(components):
    schemas = components.setdefault("schemas", {})
    if "UserInvitationToken" not in schemas:
        schemas["UserInvitationToken"] = {
            "type": "object",
            "properties": {
                "createdAt": {"type": "string", "format": "date-time"},
            },
            "required": ["createdAt"],
        }
    if "UserInvitation" not in schemas:
        schemas["UserInvitation"] = {
            "type": "object",
            "properties": {
                "organizationId": {"type": "integer"},
                "email": {"type": "string", "format": "email"},
                "tokenId": {"type": "integer"},
                "createdByUserId": {"type": "integer"},
                "createdAt": {"type": "string", "format": "date-time"},
                "updatedAt": {"type": "string", "format": "date-time"},
                "token": {"$ref": "#/components/schemas/UserInvitationToken"},
            },
            "required": [
                "organizationId",
                "email",
                "tokenId",
                "createdByUserId",
                "createdAt",
                "updatedAt",
                "token",
            ],
        }
    if "UserInvitationCreateResult" not in schemas:
        schemas["UserInvitationCreateResult"] = {
            "type": "object",
            "properties": {
                "organizationId": {"type": "integer"},
                "email": {"type": "string", "format": "email"},
                "tokenId": {"type": "integer"},
                "createdByUserId": {"type": "integer"},
                "createdAt": {"type": "string", "format": "date-time"},
                "updatedAt": {"type": "string", "format": "date-time"},
            },
            "required": [
                "organizationId",
                "email",
                "tokenId",
                "createdByUserId",
                "createdAt",
                "updatedAt",
            ],
        }


def main():
    if len(sys.argv) != 2:
        print("usage: patch_swagger.py <swagger.json>", file=sys.stderr)
        return 2

    path = sys.argv[1]
    with open(path, "r", encoding="utf-8") as f:
        doc = json.load(f)

    paths = doc.setdefault("paths", {})
    if LIST_INVITATIONS_PATH not in paths:
        paths[LIST_INVITATIONS_PATH] = {}

    list_op = paths[LIST_INVITATIONS_PATH].get("get")
    if list_op is None:
        list_op = {
            "operationId": "listUserInvitations",
            "summary": "List pending user invitations",
            "description": (
                "Lists active (non-expired) invitations for the organization. "
                "This endpoint is known to exist but is not yet published in the Swagger spec."
            ),
            "tags": ["Users"],
            "security": [{"ApiToken": []}],
            "responses": {
                "200": {
                    "description": "Invitation list",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "array",
                                "items": {"$ref": "#/components/schemas/UserInvitation"},
                            }
                        }
                    },
                },
                "401": {
                    "description": "Unauthorized",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "403": {
                    "description": "Forbidden",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaForbiddenError"}
                        }
                    },
                },
                "500": {
                    "description": "Server Error",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
            },
        }
        paths[LIST_INVITATIONS_PATH]["get"] = list_op

    if DELETE_INVITATION_PATH not in paths:
        paths[DELETE_INVITATION_PATH] = {}

    delete_op = paths[DELETE_INVITATION_PATH].get("delete")
    if delete_op is None:
        delete_op = {
            "operationId": "deleteUserInvitation",
            "summary": "Delete a pending user invitation by email",
            "description": (
                "Deletes a pending invitation for the given email address. "
                "This endpoint is known to exist but is not yet published in the Swagger spec."
            ),
            "tags": ["Users"],
            "parameters": [
                {
                    "name": "email",
                    "in": "path",
                    "required": True,
                    "schema": {"type": "string", "format": "email"},
                    "description": "Email address of the invited user to remove.",
                }
            ],
            "security": [{"ApiToken": []}],
            "responses": {
                "200": {
                    "description": "Invitation deleted",
                    "content": {"text/plain": {"schema": {"type": "string"}}},
                },
                "400": {
                    "description": "Bad Request",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "401": {
                    "description": "Unauthorized",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "403": {
                    "description": "Forbidden",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaForbiddenError"}
                        }
                    },
                },
                "404": {
                    "description": "Not Found",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "500": {
                    "description": "Server Error",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
            },
        }
        paths[DELETE_INVITATION_PATH]["delete"] = delete_op

    if SEND_INVITATION_PATH not in paths:
        paths[SEND_INVITATION_PATH] = {}

    post_op = paths[SEND_INVITATION_PATH].get("post")
    if post_op is None:
        post_op = {
            "operationId": "sendUserInvitation",
            "summary": "Send an invitation to join an organization",
            "description": (
                "Sends an invitation to join the organization. "
                "This endpoint is known to exist but is not yet published in the Swagger spec."
            ),
            "tags": ["Users"],
            "parameters": [
                {
                    "name": "id",
                    "in": "path",
                    "required": True,
                    "schema": {"type": "integer"},
                    "description": "Organization ID.",
                }
            ],
            "requestBody": {
                "required": True,
                "content": {
                    "application/json": {
                        "schema": {
                            "type": "object",
                            "properties": {
                                "emails": {
                                    "type": "array",
                                    "items": {"type": "string", "format": "email"},
                                }
                            },
                            "required": ["emails"],
                        }
                    }
                },
            },
            "security": [{"ApiToken": []}],
            "responses": {
                "200": {
                    "description": "Invitation created",
                    "content": {
                        "application/json": {
                            "schema": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/components/schemas/UserInvitationCreateResult"
                                },
                            }
                        }
                    },
                },
                "400": {
                    "description": "Bad Request",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "401": {
                    "description": "Unauthorized",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "403": {
                    "description": "Forbidden",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaForbiddenError"}
                        }
                    },
                },
                "404": {
                    "description": "Not Found",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "409": {
                    "description": "Conflict",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
                "500": {
                    "description": "Server Error",
                    "content": {
                        "application/json": {
                            "schema": {"$ref": "#/components/schemas/FossaApiError"}
                        }
                    },
                },
            },
        }
        paths[SEND_INVITATION_PATH]["post"] = post_op

    # Ensure error schemas exist (in case upstream spec changes)
    components = doc.setdefault("components", {})
    ensure_object_schema(components, "FossaApiError")
    ensure_object_schema(components, "FossaForbiddenError")
    ensure_invitation_schemas(components)

    with open(path, "w", encoding="utf-8") as f:
        json.dump(doc, f, ensure_ascii=True, separators=(",", ":"))
        f.write("\n")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
