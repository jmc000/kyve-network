<!--
order: 1
-->

# Concepts

The stakers module manages all properties of the stakers except their stake
(this is done by the delegation module). It handles staker creation, metadata
(like moniker, logo, website and commission) and the joining and leaving of
pools using valaccounts.

## Code Structure

This module adheres to our global coding structure, defined [here](../../../CodeStructure.md).

## Staker
Every address can create one staker (itself). A staker has the following
metadata which can be changed at any time.
- Moniker
- Logo
- Website

Additionally, a staker can specify a commission. However, this takes 
`CommissionChangeTime` seconds of time before the change is applied.

## Valaccounts
To join a pool the user creates a valaccount for this pool.
The existence of a valaccount (for a pool) means that the staker 
is a member of the given pool and needs to comply with the protocol
in order to not get slashed. 
A valaccount consists of the poolId a valaddress and a points counter. 
The valaddress is the address of the protocol node which is allowed
to vote in favor of the staker. For certain types of misbehavior 
(e.g. being offline) a staker collects points. These are also 
stored in the valaccount.

If a staker wants to leave a pool, a queue entry must be created. After
`LeavePoolTime` seconds of time the actual leaving is performed and the
staker can stop the protocol node for the given pool. 