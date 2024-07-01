// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
include "../src/Index.dfy"

module SimpleUnionImplTest {
    import SimpleUnion
    import opened SimpleUnionTypes
    import opened Wrappers
    method{:test} TestUnion(){
        var client :- expect SimpleUnion.SimpleUnion();
        TestMyUnionInteger(client);
        TestMyUnionString(client);
        TestKnownValueUnionString(client);
    }

    method TestMyUnionInteger(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var ret :- expect client.GetUnion(GetUnionInput(union := Some(IntegerValue(100))));

        expect ret.union.Some?;
        expect ret.union.value.IntegerValue?;
        expect ret.union.value.IntegerValue == 100;
        expect ret.union.value.StringValue? == false;
    }

    method TestMyUnionString(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var ret :- expect client.GetUnion(GetUnionInput(union := Some(StringValue("TestString"))));

        expect ret.union.Some?;
        expect ret.union.value.StringValue?;
        expect ret.union.value.StringValue == "TestString";
        expect ret.union.value.IntegerValue? == false;
    }

    method TestMyUnionBoolean(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var ret :- expect client.GetUnion(GetUnionInput(union := Some(BooleanValue(true))));

        expect ret.union.Some?;
        expect ret.union.value.BooleanValue?;
        expect ret.union.value.BooleanValue == true;
        expect ret.union.value.IntegerValue? == false;
        expect ret.union.value.StringValue? == false;
    }

    method TestMyUnionBlob(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var s: seq<UInt.uint8> := [0x0, 0x1, 0x2];
        var ret :- expect client.GetUnion(GetUnionInput(union := Some(BlobValue(s))));

        expect ret.union.Some?;
        expect ret.union.value.BlobValue?;
        expect ret.union.value.BlobValue == s;
        expect ret.union.value.BooleanValue? == false;
        expect ret.union.value.IntegerValue? == false;
        expect ret.union.value.StringValue? == false;
        expect ret.union.value.DoubleValue? == false;
    }

    method TestMyUnionDouble(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var s: seq<UInt.uint8> := [0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7];
        var ret :- expect client.GetUnion(GetUnionInput(union := Some(DoubleValue(s))));

        expect ret.union.Some?;
        expect ret.union.value.DoubleValue?;
        expect ret.union.value.DoubleValue == s;
        expect ret.union.value.BlobValue? == false;
        expect ret.union.value.BooleanValue? == false;
        expect ret.union.value.IntegerValue? == false;
        expect ret.union.value.StringValue? == false;
    }

    method TestKnownValueUnionString(client: ISimpleUnionClient)
      requires client.ValidState()
      modifies client.Modifies
      ensures client.ValidState()
    {
        var ret :- expect client.GetKnownValueUnion(GetKnownValueUnionInput(union := Some(Value(10))));

        expect ret.union.Some?;
        expect ret.union.value.Value?;
        expect ret.union.value.Value == 10;
    }
}
